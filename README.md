> NOTE: This is still a work in progress 

# About 

The goals of `strest` are to ease development, test endpoints and finally stress test an HTTP/RESTful-ish api.
Basicaly you have to describe a bunch of requests in a yaml file (a play file).
This comes handy when developing a new endpoint, api, etc. When you make changes the you play the file and look at the results.
A play file could also be played in a `test` mode. In this mode request expectations are checked.

# Roadmap
- [ ] Parse and play a play file
- [ ] Assert expectations
- [ ] Stress play a play file
- [ ] Record stats from a stres play