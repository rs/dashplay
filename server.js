var http = require('http');
var fs = require('fs');

var server = http.createServer(function(req, res) {
    var path = __dirname + '/dashplay.html';
    if (req.url.indexOf('/static/') == 0) {
        path = __dirname + req.url; // very unsafe, dev only
    }
    fs.readFile(path, function(err, data) {
        res.writeHead(200);
        res.end(data);
    });
});
server.listen(8080);
