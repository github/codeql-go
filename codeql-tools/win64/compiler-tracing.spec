**/go.exe:
  invoke ${config_dir}/go-extractor.exe
  prepend --mimic
  prepend "${compiler}"
