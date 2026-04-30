#!/bin/bash

INPUT_DIR="ressources"
OUTPUT_DIR="../nbt/registries"


REGISTRIES=(
    "banner_pattern"
    "chat_type"
    "damage_type"
    "dialog"
    "dimension_type"
    "enchantment"
    "instrument"
    "jukebox_song"
    "painting_variant"
    "test_environment"
    "test_instance"
    "trim_material"
    "trim_pattern"
    "worldgen/biome"
    "cat_variant"
    "chicken_variant"
    "cow_variant"
    "frog_variant"
    "pig_variant"
    "wolf_variant"
    "wolf_sound_variant"
    "zombie_nautilus_variant"
    "timeline"
)

mkdir -p "$OUTPUT_DIR"

for reg in "${REGISTRIES[@]}"; do
    var_name=$(basename "$reg")
    
    input_path="$INPUT_DIR/$reg"


    if [ ! -d "$input_path" ]; then
        echo "Warning: Input directory $input_path does not exist."
        continue
    fi
    
    echo "Generating $var_name registry..."
    
    python3 json_nbt_gen.py "$input_path" "$OUTPUT_DIR" registries "$var_name"
done

echo "All files generated."
