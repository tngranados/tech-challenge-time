import Config from '../config/Config';

export class ApiBase {
    protected url: string;

    constructor() {
        let url = Config.serviceURL;

        if (!url) {
            throw Error('URL cannot be empty');
        }
        if (!url.endsWith('/')) {
            url += '/';
        }
        if (!url.endsWith('api/')) {
            url += 'api/';
        }
        this.url = url;
    }

    getUrl(url: string) {
        if (!url.startsWith('http://') && !url.startsWith('https://')) {
            if (url && url.startsWith('/')) {
                url = url.substr(1);
            }
            url = `${this.url}${url}`;
        }

        return url;
    }

    protected async fetch<T>(url: string, init?: RequestInit): Promise<T> {
        // Build url for current request. Use base url as root in case it does
        // not start with http or https.
        url = this.getUrl(url);

        return new Promise<T>((resolve, reject) => {
            fetch(url, init)
                .then(res => {
                    if (res.status < 200 || res.status >= 300) {
                        res.json()
                            .then(data => reject(data))
                            .catch(e => reject(`Error: ${res.statusText}`));
                    } else {
                        res.json()
                            .then(data => resolve(data))
                            .catch(e => resolve());
                    }
                })
                .catch(e => reject(e));
        });
    }
}
