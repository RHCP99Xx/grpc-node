'use strict'
const nodemailer = require("nodemailer");
async function sendEmail(destinationEmail, token){
    let isSent = false;

    let transporter = nodemailer.createTransport({
        host: 'smtp.gmail.com',
        port: 587,
        secure: false, // true for 465, false for other ports
        auth: {
            user: 'bastagameservice@gmail.com', // generated ethereal user
            pass: '30dpr4319n'  // generated ethereal password
        },
        tls:{
          rejectUnauthorized:false
        }
      });
    
      let info = await transporter.sendMail({
        from: '"OnTheWay 🛵" OnTheWayService@gmail.com', // sender address
        to: destinationEmail + ", " + destinationEmail, // list of receivers
        subject: "OnTheWay Token ✔", // Subject line
        text: "", // plain text body
        html: "<b>Estimado ususario, este su token: </b>" + token, // html body
      }).then(error=>{
          console.log(nodemailer.getTestMessageUrl(error))
      });
      /*
      if(info.messageId != null){
          isSent = true;
      }else{
          isSent = false;
      }
    */
    return isSent;
}
module.exports = {
    sendEmail,
}