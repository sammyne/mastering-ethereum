
if ($args.Count -eq 0) {
  echo 'too few arguments'
  exit
}

docker run --rm ethereum/client-go:alltools-v1.8.23 geth $args