

import fetch from 'node-fetch'

class MapService {
    fetch() {
        return fetch('http://localhost:9000/api/map')
            .then((res) => {
                return res.json();
            });
    }
}

export default MapService
