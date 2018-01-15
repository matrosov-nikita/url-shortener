const express = require('express');
const bodyParser = require('body-parser');
const rp = require('request-promise');
const url = require('url');
const app = express();

const genURLOptions = (uri, params) => {
  return {
    method: 'POST',
    uri: `${process.env.API_URL}/urlshortener/URLHandler/${uri}`,
    json: true,
    body: params
  };
};

app.use(express.static('public'));
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());


app.post('/encode', (req, res, next) => {
  rp(genURLOptions('encode', req.body)).then(resp => {
    res.send(resp);
  }).catch(err => {
    next({
      code: 500,
      message: `could not encode: ${err}`
    });
  });
});

app.get('/:hash', (req, res) => {
  rp(genURLOptions('decode', { 'url': req.params.hash })).then(resp => {
    return url.parse(resp.originUrl).protocol ? resp.originUrl: `http://${resp.originUrl}`;
  }).catch(err => {
    return '/';
  }).then(URL => res.redirect(URL));
});

app.use((err, req, res) => {
  res.status(err.code).send(err.message);
});

app.listen(3000);