https://appspector.com/blog/how-to-improve-messagepack-javascript-parsing-speed-by-2-6-times

msgpack is slower for some things and faster for others in Dart and JS, it seems like wins over JSON aren't really worth the loss of readability.

Go msgpack libraries are 10x faster than the standard lib json. I wonder if easyjson or something similar could provide a good middle-of-the-road option here between performance and readability.s