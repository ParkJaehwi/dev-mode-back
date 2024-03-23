const express = require('express');
const app = express();
let path = require('path');
let sdk = require('./sdk');

const PORT = 8001;
const HOST = '0.0.0.0';
app.use(express.json());
app.use(express.urlencoded({ extended: true }))

app.post('/CreateIdentity', function (req, res) {
   let name = req.body.name;
   let gender = req.body.gender;
   let dob = req.body.dob;
   let contact = req.body.contact;
   let idnumber = req.body.idnumber;
   let args = [name, gender, dob, contact, idnumber];
   sdk.send(false, 'CreateIdentity', args, res);
});


app.get('/QueryIdentity', function (req, res) {
   
   let idnumber = req.query.idnumber;
   let args = [idnumber];
   sdk.send(false, 'QueryIdentity', args, res);
});


app.use(express.static(path.join(__dirname, '../client')));
app.listen(PORT, HOST);
    console.log(`Running on http://${HOST}:${PORT}`);
