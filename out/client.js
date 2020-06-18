import { fetch } from 'fetch-h2';
import { encode, decode } from 'msgpack-lite';

const client = (host, https = true) => {
    const scheme = https ? "https" : "http";

    return {
        createUser: async () => {
            const url = `${scheme}://${host}/rpc/v1/createUser`;
            const payload = {};

            const response = await fetch(url, {
                method: 'OPTIONS',
                // mode: 'cors', // no-cors, *cors, same-origin
                // cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
                // credentials: 'same-origin', // include, *same-origin, omit
                // headers: {
                //     'Content-Type': 'application/x-msgpack'
                // },
                // redirect: 'follow', // manual, *follow, error
                // referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
                body: JSON.stringify(payload),
            });

            return await response.json();
        },
        archiveUser: (id) => {},
        createCocktail: async () => {
            const url = `${scheme}://${host}/rpc/v1/createCocktail`;
            const payload = {};

            const response = await fetch(url, {
                method: 'OPTIONS',
                // mode: 'cors', // no-cors, *cors, same-origin
                // cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
                // credentials: 'same-origin', // include, *same-origin, omit
                // headers: {
                //     'Content-Type': 'application/x-msgpack'
                // },
                // redirect: 'follow', // manual, *follow, error
                // referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
                body: JSON.stringify(payload),
            });

            return await response.json();
        },
        archiveCocktail: (id) => {},
    };
};