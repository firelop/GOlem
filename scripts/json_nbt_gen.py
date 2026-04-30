import os, os.path, json, sys

def to_camel_case(snake_str):
    components = snake_str.split('_')
    return "".join(x.title() for x in components)

def exclude_mc(s):
    s = s.replace("/", "_")
    if s.startswith("minecraft:"):
        return s[len("minecraft:"):]
    return s

# ---------------------------------------------------------------------------
# Schema merging
#
# merge_dicts builds a "merged schema" across all entries in a registry.
# When the same key maps to structurally incompatible values across entries
# (different dict key-sets, or different primitive types), it is marked as
# the sentinel string "conflict_any", which causes get_go_type to emit `any`
# and generate_go_value to emit a map[string]any literal.
# ---------------------------------------------------------------------------

def merge_dicts(d1, d2):
    """Merge two schema-dicts. Marks divergent fields as 'conflict_any'."""
    res = dict(d1)
    for k, v in d2.items():
        if k not in res:
            res[k] = v
            continue
        existing = res[k]
        if existing == "conflict_any":
            continue
        if isinstance(existing, dict) and isinstance(v, dict):
            if set(existing.keys()) != set(v.keys()):
                res[k] = "conflict_any"
            else:
                res[k] = merge_dicts(existing, v)
        elif isinstance(existing, list) and isinstance(v, list):
            # Empty lists are type-less — prefer whichever side is non-empty
            if not existing:
                res[k] = v  # take the non-empty (or equally empty) side
            elif not v:
                pass  # keep existing non-empty schema
            elif isinstance(existing[0], dict) and isinstance(v[0], dict):
                # Merge list-of-dict schemas; if key-sets are incompatible, mark conflict
                merged_item = merge_list_item_schema(existing + v)
                if merged_item is None:
                    res[k] = "conflict_any"
                else:
                    res[k] = [merged_item]
            # primitive lists: keep as-is (same type, nothing to reconcile)
        elif type(existing) != type(v):
            res[k] = "conflict_any"
        # same type, same value shape: keep existing
    return res

def merge_list_item_schema(items):
    """
    Merge a list of dicts into a single schema dict.
    Returns None if any two items have different key-sets (unstable structure).
    """
    dict_items = [x for x in items if isinstance(x, dict)]
    if not dict_items:
        return {}
    first_keys = set(dict_items[0].keys())
    for item in dict_items[1:]:
        if set(item.keys()) != first_keys:
            return None
    result = dict(dict_items[0])
    for item in dict_items[1:]:
        result = merge_dicts(result, item)
    return result

# ---------------------------------------------------------------------------
# Type inference
#
# get_go_type returns the Go type string for a schema value.
# ---------------------------------------------------------------------------

def get_go_type(schema, indent=""):
    if schema == "conflict_any":
        return "any"
    if isinstance(schema, dict):
        lines = "struct {\n"
        for k, v in sorted(schema.items()):
            lines += f"{indent}\t{to_camel_case(exclude_mc(k))} {get_go_type(v, indent + chr(9))} `nbt:\"{k}\"`\n"
        lines += indent + "}"
        return lines
    if isinstance(schema, list):
        if not schema:
            return "[]any"
        elem = schema[0]
        return f"[]{get_go_type(elem, indent)}"
    if isinstance(schema, bool):
        return "bool"
    if isinstance(schema, int):
        return "int32"
    if isinstance(schema, float):
        return "float32"
    if isinstance(schema, str):
        return "string"
    return "any"

# ---------------------------------------------------------------------------
# Value emission
#
# generate_go_value emits a Go literal for a concrete data value,
# using `schema` to govern type declarations. When schema == "conflict_any",
# the value is emitted as a map[string]any literal.
# ---------------------------------------------------------------------------

def go_scalar(data):
    if isinstance(data, bool):
        return "true" if data else "false"
    if isinstance(data, int):
        return f"int32({data})"
    if isinstance(data, float):
        return f"float32({data})"
    if isinstance(data, str):
        escaped = data.replace("\\", "\\\\").replace('"', '\\"')
        return f'"{escaped}"'
    if data is None:
        return "nil"
    return str(data)

def generate_go_value(data, schema, indent="", is_list_item=False):
    """
    Emit a Go literal for `data`, driven by `schema`.

    - schema == "conflict_any"  → map[string]any literal
    - schema is a dict          → typed struct literal
    - schema is a list          → typed slice literal
    - schema is a scalar        → typed scalar literal
    """
    # Conflict: emit as map[string]any so the NBT encoder handles it dynamically
    if schema == "conflict_any":
        if isinstance(data, dict):
            return emit_map_any(data, indent)
        if isinstance(data, list):
            return emit_slice_any(data, indent)
        return go_scalar(data)

    if isinstance(data, dict):
        if is_list_item:
            # Inside a typed slice: just emit {Field: value, ...}
            res = "{\n"
            for k, v in sorted(data.items()):
                sub_schema = schema.get(k, v) if isinstance(schema, dict) else v
                res += f"{indent}\t{to_camel_case(exclude_mc(k))}: {generate_go_value(v, sub_schema, indent + chr(9))},\n"
            res += indent + "}"
            return res
        else:
            # Top-level or nested struct: emit type declaration + value
            res = "struct {\n"
            for k, v in sorted(data.items()):
                sub_schema = schema.get(k, v) if isinstance(schema, dict) else v
                res += f"{indent}\t{to_camel_case(exclude_mc(k))} {get_go_type(sub_schema, indent + chr(9))} `nbt:\"{k}\"`\n"
            res += indent + "}{\n"
            for k, v in sorted(data.items()):
                sub_schema = schema.get(k, v) if isinstance(schema, dict) else v
                res += f"{indent}\t{to_camel_case(exclude_mc(k))}: {generate_go_value(v, sub_schema, indent + chr(9))},\n"
            res += indent + "}"
            return res

    if isinstance(data, list):
        if not data:
            # Emit a typed empty slice matching the schema
            if isinstance(schema, list) and schema:
                elem_schema = schema[0]
                return f"[]{get_go_type(elem_schema, indent)}{{}}"
            return "[]any{}"
        # Determine element schema
        if isinstance(schema, list) and schema:
            elem_schema = schema[0]
        else:
            elem_schema = data[0]  # fallback: infer from first element

        if elem_schema == "conflict_any" or (isinstance(elem_schema, dict) and
                merge_list_item_schema([x for x in data if isinstance(x, dict)]) is None):
            return emit_slice_any(data, indent)

        # If elements are themselves lists (e.g. [][]any), use emit_any_value
        # so inner lists are always typed as []any (assignable to any slots)
        if isinstance(elem_schema, list):
            elem_type = get_go_type(elem_schema, indent)
            res = f"[]{elem_type}{{\n"
            for item in data:
                res += f"{indent}\t{emit_any_value(item, indent + chr(9))},\n"
            res += indent + "}"
            return res

        elem_type = get_go_type(elem_schema, indent)
        res = f"[]{elem_type}{{\n"
        for item in data:
            res += f"{indent}\t{generate_go_value(item, elem_schema, indent + chr(9), is_list_item=isinstance(item, dict))},\n"
        res += indent + "}"
        return res

    return go_scalar(data)

def emit_map_any(data, indent):
    """Emit a map[string]any{...} literal."""
    res = "map[string]any{\n"
    for k, v in sorted(data.items()):
        escaped_k = k.replace("\\", "\\\\").replace('"', '\\"')
        res += f'{indent}\t"{escaped_k}": {emit_any_value(v, indent + chr(9))},\n'
    res += indent + "}"
    return res

def emit_slice_any(data, indent):
    """Emit a []any{...} literal. All nested lists are also []any so they
    are assignable to any-typed slots without type mismatch."""
    res = "[]any{\n"
    for item in data:
        res += f"{indent}\t{emit_any_value(item, indent + chr(9))},\n"
    res += indent + "}"
    return res

def emit_any_value(data, indent):
    """Emit a value for use inside map[string]any or []any.
    Lists are always emitted as []any regardless of element type,
    so they stay compatible with any-typed parent fields."""
    if isinstance(data, dict):
        return emit_map_any(data, indent)
    if isinstance(data, list):
        # Always []any so it can be assigned to any-typed parent
        return emit_slice_any(data, indent)
    return go_scalar(data)

# ---------------------------------------------------------------------------
# Main generation logic (two-pass)
# ---------------------------------------------------------------------------

def generate_nbt_json(input_dir, output_dir, package_name, var_name):
    os.makedirs(output_dir, exist_ok=True)

    files = sorted(f for f in os.listdir(input_dir) if f.endswith(".json"))

    # Pass 1: load all entries and build a merged cross-entry schema.
    # Fields whose structure varies across entries become "conflict_any".
    entries = []
    master_schema = {}
    for filename in files:
        with open(os.path.join(input_dir, filename), "r") as f:
            data = json.load(f)
        identifier = filename[:-len(".json")]
        entries.append((identifier, data))
        master_schema = merge_dicts(master_schema, data)

    # Pass 2: emit Go code, using master_schema to drive type declarations.
    with open(os.path.join(output_dir, var_name + ".go"), "w") as go_file:
        go_file.write(f"package {package_name}\n\n")
        go_file.write(f"var {to_camel_case(var_name)} = map[string]any{{\n")
        for identifier, data in entries:
            go_file.write(f'\t"{identifier}": ')
            go_file.write(generate_go_value(data, master_schema, "\t"))
            go_file.write(",\n")
        go_file.write("}\n")

if __name__ == "__main__":
    if len(sys.argv) < 5:
        print("Usage: python json_nbt_gen.py <input_dir> <output_dir> <package_name> <var_name>")
        sys.exit(1)
    generate_nbt_json(sys.argv[1], sys.argv[2], sys.argv[3], sys.argv[4])
