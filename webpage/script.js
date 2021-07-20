function Clock() {
  var now = new Date();
  var hour = now.getHours();
  var min  = ('0'+now.getMinutes()).slice(-2);
  var sec  = ('0'+now.getSeconds()).slice(-2);

  var msg = "現在時刻 " + hour + ":" + min + ":" + sec;
  document.getElementById("TimeArea").innerHTML = msg;
}

setInterval('Clock()',1000);

function CurtainMove(status) {
  var msg2 = "現在の状態は" + status;
  document.getElementById("Status").innerHTML = msg2;
  console.log("カーテンの操作が行われた");
}
