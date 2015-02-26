# go1billion
The repository has information necessary to run the Golang Count to One Billlllllllionnnnnn test.

## Live Version on Run.Pivotal.Io
[http://go1billion.cfapps.io/](http://go1billion.cfapps.io/)



## Push to Cloud Foundry!
Modify it with a custom name of course..

    git clone https://github.com/clintonskitson/go1billion
    cd go1billion
    ./cf push go1billion

## Run in a Docker container
The following command will run the Docker container.

    docker run -tid -p 8080:8080 clintonskitson/go1billion

If you are running boot2docker make sure to reference the IP of the VM which is the endpoint you need to target with ```boot2docker ip```.

## Cross Compile
The easiest way to cross compile the app is to leverage a Docker container.  Below is a command that you can run to *generate binaries* across multiple platforms.

    go get github.com/clintonskitson/go1billion
    docker run --rm -it -v $GOPATH:/go -w /go/src/github.com/clintonskitson/go1billion golang:1.4-cross make release

This should generate the following binaries across referenced platforms.

    -rwxr-xr-x  1 clintonkitson  staff  4919548 Feb 25 19:22 go1billion-Darwin-i386
    -rwxr-xr-x  1 clintonkitson  staff  6126368 Feb 25 19:22 go1billion-Darwin-x86_64
    -rwxr-xr-x  1 clintonkitson  staff  6027616 Feb 25 19:22 go1billion-FreeBSD-amd64
    -rwxr-xr-x  1 clintonkitson  staff  4823672 Feb 25 19:22 go1billion-FreeBSD-i386
    -rwxr-xr-x  1 clintonkitson  staff  4879928 Feb 25 19:22 go1billion-Linux-armv6l
    -rwxr-xr-x  1 clintonkitson  staff  4879928 Feb 25 19:22 go1billion-Linux-armv7l
    -rwxr-xr-x  1 clintonkitson  staff  4854328 Feb 25 19:22 go1billion-Linux-i386
    -rwxr-xr-x  1 clintonkitson  staff  4854328 Feb 25 19:22 go1billion-Linux-i686
    -rwxr-xr-x  1 clintonkitson  staff  4221920 Feb 25 19:22 go1billion-Linux-static
    -rwxr-xr-x  1 clintonkitson  staff  6019784 Feb 25 19:22 go1billion-Linux-x86_64
    -rwxr-xr-x  1 clintonkitson  staff  4930048 Feb 25 19:22 go1billion.exe

## Run Binary
You can choose any of the proper binaries listed above.  If you are running it form OS X you would likely leverage this command ```./release/go1billion-Darwin-x86_64```.

## Parameter
There is a single parameter currently that can be passed.  The ```numcpu``` parameter will specify how many Go routines that should run in parallel.  These routines will split the 1 billion iterations.

## API returns
The following is what you can expect from the API.

```http://go1billion.cfapps.io/?numcpu=1```

By running this request, or a request that you require, it will generate the JSON response below.  

    {"InstanceNumCPU":4,"RequestNumCPU":1,"CountTo":1e+09,"CountToPer":1000000000,"StartingTime":"2015-02-26T03:06:23.204825031Z","EndingTime":"0001-01-01T00:00:00Z","Duration":430786666,"Durationms":430}
