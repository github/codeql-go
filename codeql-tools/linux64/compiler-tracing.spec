**/go:
  invoke ${config_dir}/go-extractor
  prepend --mimic
  prepend "${compiler}"
