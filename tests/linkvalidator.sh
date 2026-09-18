linkvalidator --root-dir ../. --exclude-domains "github.com" --hide-warnings > output.txt || EXIT_CODE=$?
cat output.txt && exit ${EXIT_CODE}
                
                
                
              
