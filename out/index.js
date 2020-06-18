const { fetch } = require('fetch-h2');
const msgpack = require('msgpack-lite');

fetch("http://localhost:5000/msg", {
    method: 'POST',
    // mode: 'cors', // no-cors, *cors, same-origin
    // cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    // credentials: 'same-origin', // include, *same-origin, omit
    // headers: {
    //     'Content-Type': 'application/x-msgpack'
    // },
    // redirect: 'follow', // manual, *follow, error
    // referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: msgpack.encode({
        message: "hello world",
    }),
})
    // .then((res) => decode(res.body))
    .then((res) => res.arrayBuffer())
    .then((res) => msgpack.decode(new Uint8Array(res)))
    .then((res) => console.log(res));
