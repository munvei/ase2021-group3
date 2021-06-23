 <!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>home</title>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.min.js"></script>
    <link rel="stylesheet"href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
  </head>
  <body class="white black-text container flow-text">
    <h1 class="center">ACOS</h1>
    <center style="font-size:100%" id="TimeArea"></center>
    <div>アラーム設定</div>
    <?php 
      if(isset($_POST['alarm_time'])){
        include('create_file.php');
      }
      $file = "test.txt";
      $line = file_get_contents($file);
      echo "現在の設定時刻($line)"
    ?>
    <form action="#" method="POST">
      <p>アラーム時刻:<br>
      <input type="time" name="alarm_time"></p>
      <p><input type="submit" value="決定"></p>
    </form>
    <div>カーテン操作</div>
    <input type="button" value="操作" onclick="CurtainMove();">
    <script type="text/javascript" src="./script.js"></script>
  </body>
</html>

