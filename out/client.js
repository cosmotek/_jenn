import { fetch } from 'fetch-h2';
import { encode, decode } from 'msgpack-lite';

const client = (host, https = true) => {
    const scheme = https ? "https" : "http";
    
    return {
        createUser: async () => {
            const url = `${scheme}://${host}/api/v1/rpc/createUser`;
            const payload = {};

            const response = await fetch(url, {
                method: 'OPTIONS',
                mode: 'cors', // no-cors, *cors, same-origin
                cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
                credentials: 'same-origin', // include, *same-origin, omit
                headers: {
                    'Content-Type': 'application/json'
                },
                redirect: 'follow', // manual, *follow, error
                referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
                body: encode(payload),
            });

            return decode(response.body);
        },
        archiveUser: (id) => {},
        createCocktail: async () => {
            const url = `${scheme}://${host}/api/v1/rpc/createCocktail`;
            const payload = {};

            const response = await fetch(url, {
                method: 'OPTIONS',
                mode: 'cors', // no-cors, *cors, same-origin
                cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
                credentials: 'same-origin', // include, *same-origin, omit
                headers: {
                    'Content-Type': 'application/json'
                },
                redirect: 'follow', // manual, *follow, error
                referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
                body: encode(payload),
            });

            return decode(response.body);
        },
        archiveCocktail: (id) => {},
    };
};