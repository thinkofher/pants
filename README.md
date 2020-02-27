pants
-------

Pants is a simple service for link shortening, builded with `go`, `elm` and love. Created only for educational purposes. It provides beautiful UI expierence and simple REST API. It's really easy to build and run it!

You can check sample running application [here](http://short.beniamindudek.xyz). Please have mercy when using it.

Requirements
------------

Basically you just need to install `go` and `elm` compiler. But if you're not using some decent unix-like operatiny system you will also need to install `GNU/make`.


Installation
------------

Clone this repo, enter the directory and run:

    $ cd app
    $ make app
    $ ./pants

Voila! You are running `pants` on port `8080`. Contragulations!

Containers
----------

You can easly deploy this application with provided Dockerfiles for [app](app/Dockerfile) and [view](view/Dockerfile), so you don't have to install all the dependencies and bundlers directly on your machine.

`pants` also ships with [docker-compose](docker-compose.yml) file, so you can easly hack on it with:

    $ docker-compose up -d

or

    $ podman-compose up -d

More
----

Check out [API](app/README.md) specification for further informations.

Contribution
------------

Fork it, create your own branch, hack it and pull it.
