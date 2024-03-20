#!/bin/bash

# Set up directories relative to the current directory
workdir="cache"
localesdir="locales"
ignored="^(vendor|resources|cache|catalog|$workdir)"

# Function to extract messages
extract_messages() {
  if [ -e "$workdir" ]; then
    rm -rf "$workdir"
  fi
  mkdir "$workdir"

  # The following flags are:

  # -c extract-strings: This flag specifies the action to perform, which in this case is to extract translation strings from the source code.
  # -po: This flag indicates that the output format should be .po files.
  # -d: This flag specifies the directory to search for Go source code files. Here, . represents the current directory.
  # r: This flag tells the tool to search for Go source code files recursively within the specified directory.
  # -o "$workdir": This flag specifies the output directory where the extracted translation strings will be saved. The $workdir variable contains the path to the working directory, where the extracted strings will be stored.
  # —ignore-regexp "$ignored": This flag specifies a regular expression pattern to ignore certain files or directories during the extraction process. The $ignored variable contains the pattern for files or directories to ignore.
  # —output-match-package: This flag indicates that the extracted translation strings should be grouped by package.
  i18n4go -c extract-strings -v --po -d . -r -o "$workdir" --ignore-regexp "$ignored" -output-match-package
}

# Function to merge messages for all locales into .pot files
merge_messages() {
  # Merge all .po files into a single .po file
  msgcat --use-first "$workdir"/**/*.po -o "$workdir/merged_messages.po"

  # Add charset specification to the merged .po file
  echo 'msgid ""' > "$workdir/merged_messages_with_charset.po"
  echo 'msgstr ""' >> "$workdir/merged_messages_with_charset.po"
  echo '"Content-Type: text/plain; charset=UTF-8\n"' >> "$workdir/merged_messages_with_charset.po"
  cat "$workdir/merged_messages.po" >> "$workdir/merged_messages_with_charset.po"

  # Iterate over each subdirectory in the locales directory
  for locale_dir in "$localesdir"/*; do
    # Check if it's a directory
    if [ -d "$locale_dir" ]; then
      # Extract the locale from the directory name
      locale=$(basename "$locale_dir")
      # Check if there are any .po files in the LC_MESSAGES directory of the locale
      if compgen -G "$locale_dir/LC_MESSAGES/default.po" > /dev/null; then
        # Merge the combined .po file into the individual locale .po file without overriding existing translations
        msgcat --use-first "$locale_dir/LC_MESSAGES/default.po" "$workdir/merged_messages_with_charset.po" -o "$locale_dir/LC_MESSAGES/default_merged.po"
        # Move the merged .po file to replace the original default.po file
        mv "$locale_dir/LC_MESSAGES/default_merged.po" "$locale_dir/LC_MESSAGES/default.po"
      else
        echo "No default.po file found in $locale_dir/LC_MESSAGES. Skipping merging for $locale."
      fi
    fi
  done

  # Clean up temporary files
  rm -f "$workdir/merged_messages.po" "$workdir/merged_messages_with_charset.po"
}

# Main script
if [ $# -gt 0 ]; then
  if [ "$1" = "--extract" ]; then
    extract_messages
    merge_messages
  fi
fi