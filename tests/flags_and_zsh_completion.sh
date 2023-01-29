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
 
_test() { 
  # compadd -X "flags:" "--verbose" "-d" "--debug" "-n" "--name"  #can't use "-v"
  # _arguments '-s[sort output]' '--l[long output]' '-l[long output]' #more that
  # _arguments '-s[sort output]'


  #  _arguments '-v[sort output]:filename:_files' # arg taking file
  #  _arguments '-f[flags]:toto:(ITEM1 ITEM2)' # predefined value
  #  _arguments '-x[xoxo]:youpi:()' # mandatory value
_arguments '-f[provide a file]:filename:->files' '-p[predefined]:set:->predef'
case "$state" in
    files)
        _files 
        ;;
    predef)
        _values'set' a b c d e #only one value
        _values -s , 'set' a b c d e # list
        ;;
esac
}  
   
compdef _test test


#https://github.com/zsh-users/zsh-completions/blob/master/zsh-completions-howto.org