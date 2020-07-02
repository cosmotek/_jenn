import 'dart:async';
import 'dart:convert';
import 'dart:io';

class Client {
  String _host;
  int _port;

  HttpClient _client;

  Client(String host, int port) {
    _host = host;
    _port = port;
    _client = HttpClient();
  }

  Future<dynamic> createUser() async {
    var req = await _client.open('OPTIONS', _host, _port, '/rpc/v1/createUser');
    
    // set headers
    req.headers.set('content-type', 'application/json');
    req.headers.set('accepts', 'application/json');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      return body;
    }
  }

  Future<dynamic> streamUser() async {
    var req = await _client.open('GET', _host, _port, '/rpc/v1/streamUser');
    
    // set headers
    req.headers.set('Content-Type', 'application/json');
    req.headers.set('Accepts', 'text/event-stream');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      // return body;
      print(body);
    }
  }
  Future<dynamic> archiveUser() async {
    var req = await _client.open('OPTIONS', _host, _port, '/rpc/v1/archiveUser');
    
    // set headers
    req.headers.set('Content-Type', 'application/json');
    req.headers.set('Accepts', 'application/json');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      return body;
    }
  }

  Future<dynamic> createCocktail() async {
    var req = await _client.open('OPTIONS', _host, _port, '/rpc/v1/createCocktail');
    
    // set headers
    req.headers.set('content-type', 'application/json');
    req.headers.set('accepts', 'application/json');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      return body;
    }
  }

  Future<dynamic> streamCocktail() async {
    var req = await _client.open('GET', _host, _port, '/rpc/v1/streamCocktail');
    
    // set headers
    req.headers.set('Content-Type', 'application/json');
    req.headers.set('Accepts', 'text/event-stream');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      // return body;
      print(body);
    }
  }
  Future<dynamic> archiveCocktail() async {
    var req = await _client.open('OPTIONS', _host, _port, '/rpc/v1/archiveCocktail');
    
    // set headers
    req.headers.set('Content-Type', 'application/json');
    req.headers.set('Accepts', 'application/json');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      return body;
    }
  }

  Future<dynamic> createBeverage() async {
    var req = await _client.open('OPTIONS', _host, _port, '/rpc/v1/createBeverage');
    
    // set headers
    req.headers.set('content-type', 'application/json');
    req.headers.set('accepts', 'application/json');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      return body;
    }
  }

  Future<dynamic> streamBeverage() async {
    var req = await _client.open('GET', _host, _port, '/rpc/v1/streamBeverage');
    
    // set headers
    req.headers.set('Content-Type', 'application/json');
    req.headers.set('Accepts', 'text/event-stream');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      // return body;
      print(body);
    }
  }
  Future<dynamic> archiveBeverage() async {
    var req = await _client.open('OPTIONS', _host, _port, '/rpc/v1/archiveBeverage');
    
    // set headers
    req.headers.set('Content-Type', 'application/json');
    req.headers.set('Accepts', 'application/json');

    // add body
    req.add(utf8.encode(json.encode({})));

    var res = await req.close();
    if (res.statusCode != 200) {
      throw 'Its not all good... ${res.statusCode}';
    }

    await for (var body in res.transform(utf8.decoder)) {
      return body;
    }
  }

}

void main() {
  Client("localhost", 5000).streamUser();
}