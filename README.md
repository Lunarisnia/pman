# PMAN - Password Manager
Password manager CLI app written in golang.

# How To Use
1. Download the latest release
2. Extract the program in your desired location
3. Add the location to env path
4. Change the content of pman-key.txt to your desired password
5. Open any terminal/powershell then type `pman` to test if there is any error

# Documentation
## Command
- list  
    Show the list of your saved passwords  
- generate -n PasswordName  
    Generate a new password and save it  
- show -i PasswordID  
    Reveal a single password for use


## Flags
- generate  
    - --name, -n : name; A name for your password (required)  
    - --length, -n : password length; The length of your generated password default to 16  
- show
    - --id, -i : Password ID; The id of your password show from the list command (required)  


# Build your own
1. Clone this repo
2. run `go get`
3. build the program with `go build .`