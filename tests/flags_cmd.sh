
test(){
VERBOSE=false
DEBUG=false
NAME=
while true; do
  case "$1" in
    -v | --verbose ) VERBOSE=true; shift ;;
    -d | --debug ) DEBUG=true; shift ;;
    -n | --name ) NAME="$2"; shift 2 ;;
    -- ) shift; break ;;
    * ) break ;;
  esac
done


echo $VERBOSE
if [ -z ${NAME} ]; then echo "No name"; else echo "name is '$NAME'"; fi
}