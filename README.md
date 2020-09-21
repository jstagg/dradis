# dradis
Go http front-end for a Redis back-end.

The back-end is in "redis-back" and the front-end is "redis-front." The data is provided in Redis format via a simple SQL statement. It's imported when the container is built. Each successive change in data assumes the container will get a new version, though Redis will happily accept updates while running.

If you don't care about how the containers got built (and let's be honest, it ain't fancy, so you might not) and just want to run them, get that docker-compose.yml file and:
docker-compose up -d

The old Windows batch files for running each one (and attaching to a Docker network) are still in there should you care.

One caveat! I didn't parameterize the name of "dradis-back" in the Go program hosted in "dradis-front." So should you want to run dradis-back as a different container name, be sure to modify main.go. One day, I'll parameterize it. Doing it just wasn't core to my PoC's use case.

If the PoC use case is of any interest, it's a mini-project from work for a "decomposing a larger service" approach. Any enterprise application is apt to need to validate data on input. Ideally, pushing these activities to the edge where the microservices could be consumed by customers (or at least adge processes) provides those edge systems with a convenient way to ask "is this valid input?" The output of the call contains some data relevant to the next steps in processing. That data is serialized into a pipe-delimited string since this is readily understandable to the potential users of the microservice.

If it seems overly simple, it is. Debugging is twice as hard as building. Therefore, if we use ALL of our wits during building, we've literally outsmarted ourselves for debugging. Keep it overly simple, and debugging is much, much easier.