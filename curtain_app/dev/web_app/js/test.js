document.addEventListener('DOMContentLoaded', () => {
    let loc = window.location;
    let url = 'ws:';
    if (loc.protocol === 'https:') {
        url = 'wss:';
    }
    url += '//' + loc.host;
    url += loc.pathname + '/ws';

    const ws = new WebSocket(url)
    ws.onopen = () => {
        console.log('Connected')
    }

    ws.onmessage = (evt) => {
        let out = document.getElementById('output');
        out.innerHTML += evt.data + '<br>';
    }

    const btn = document.querySelector('.btn')
    btn.addEventListener('click', () => {
        ws.send(document.getElementById('input').value)
    })
});
