> NOTE: This is still a work in progress 

# About 

The goals of `strest` are to ease development and testing of HTTP/REST(-ish) APIs.

Basicaly you have to describe a bunch of requests in a yaml file (a play file).
Then you make changes the your API, play the play and look at the results.

When v1 is ready it should have the following general features:

- [ ] Play a bunch of requests described in a play file (the `play` command)
- [ ] Play a play file and assert expectation (the `test` command)
- [ ] Stress play a play file and gather some stats (the `stress` command)
