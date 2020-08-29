# Set GOPATH variable every time using terminal as we have multiple GOPATH locations
# These commands will set GOPATH and PATH vriables temporarily and will vanish when you close terminal
# These values for these variables are set permanently in ~/.bashrc file
export GOPATH=/home/blockdigest/Documents/Udemy/GolangTutorialsUdemy
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin

====================================================================================

# Create a new repository on the command line
echo "# tryingagain" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M master
git remote add origin https://github.com/AayushWork/tryingagain.git
git push -u origin master

====================================================================================

# When adding new file or a change to github
git add --all
git commit -m "some relevant msg"
git push

# or you can club add and commit together
git add --all && git commit -m "XXX"
git push

# To add versioninig tags
git tag v1.0.0
git push --tags

# check pushed tags by command:
git tag --list

# Branch out to a new version
git checkout -b v2
====================================================================================

# To remove username and password entering by generating ssh key
ssh-keygen -t rsa
cat /home/blockdigest/.ssh/id_rsa.pub

# Copy content and create a new SSH key on Github
# Add the public key content there and create Key

# Now you can run the git clone command by copying the SSH url gtom github
# Below command will create the folder exactly like on github without asking for password
git clone git@github.com:AayushWork/tryingagain.git

====================================================================================