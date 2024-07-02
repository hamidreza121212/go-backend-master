sudo apt update
sudo apt install curl unzip

# Download the latest release
curl -Lo sqlc.zip https://github.com/kyleconroy/sqlc/releases/latest/download/sqlc_linux_amd64.zip

# Unzip the downloaded file
unzip sqlc.zip

# Move the binary to a directory in your PATH
sudo mv sqlc /usr/local/bin/

# Clean up
rm sqlc.zip

