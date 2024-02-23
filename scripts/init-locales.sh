#!/bin/bash

# List of locales
locales=("en_US" "el_GR" "ar_SA")

# Path to the template .pot file
template_file="locales/default.po"

# Loop through each locale
for locale in "${locales[@]}"; do
    # Output directory for the .po file
    output_dir="locales/$locale/LC_MESSAGES"

    # Create the output directory if it doesn't exist
    mkdir -p "$output_dir"

    # Output file path for the .po file
    output_file="$output_dir/default.po"

    # Call msginit to initialize the .po file
    msginit --no-translator -o "$output_file" -l "$locale" -i "$template_file"

    # Print status message
    echo "Initialized $output_file"
done