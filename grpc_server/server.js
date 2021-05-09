const PROTO_PATH = __dirname + '/proto/services.proto';

const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const _ = require('lodash');

const sendEmail = require('./EmailSender/email_sender');

const nodemailer = require('nodemailer');


let packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
});

let email_token_proto = grpc.loadPackageDefinition(packageDefinition).EmailToken;

function getEmailAndToken(call, callback){
    let isSent = sendEmail.sendEmail(call.request.email, call.request.token)
    callback(null, {response: isSent});
}


function main() {
    let server = new grpc.Server();
    server.addService(email_token_proto.EmailToken.service, {getEmailAndToken: getEmailAndToken});
    server.bindAsync('0.0.0.0:4500', grpc.ServerCredentials.createInsecure(), () =>{
        server.start();
    });
  }

  main();