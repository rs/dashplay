var http = require('http');
var fs = require('fs');

var server = http.createServer(function(req, res) {
    var path = __dirname + '/web/';

    if (req.url.indexOf('/static/') == 0) {
        path += req.url; // very unsafe, dev only
    } else {
        path += 'dashplay.html';
    }

    fs.readFile(path, function(err, data) {
        res.writeHead(200);
        res.end(data);
    });
});

server.listen(8080);
