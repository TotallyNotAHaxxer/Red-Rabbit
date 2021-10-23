sudo apt-get install docker 
sudo apt-get update 
git clone  https://github.com/mackwic/colored.git
cd colored 
docker build -t colored_image .
docker run --rm -it -v "$PWD":/src -u `id -u`:`id -g` colored_image /bin/bash -c "cargo build"
echo 'testing'
docker run --rm -it -v "$PWD":/src -u `id -u`:`id -g` colored_image /bin/bash -c "cargo test"

