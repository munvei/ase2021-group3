 <!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>home</title>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.min.js"></script>
    <link rel="stylesheet"href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
    <link rel="stylesheet"href="style.css">
  </head>
  <body class="white black-text container flow-text">
    <header class="blue white-text">
    <h1 class="center">ACOS 管理ページ</h1>
    </header>
    <center style="font-size:120%" id="TimeArea"></center>
    <?php 
      if(isset($_POST['alarm_time'])){
        include('create_file.php');
      }
      $file = "test.txt";
      $line = file_get_contents($file);
      echo "現在の設定時刻($line)<br/>"
    ?>
    <div style="font-size:100%" id="Status" class="blue item">現在の状態は</div>
    <div>カーテン操作</div>
    <input class="blue left-btn" type="button" value="開" onclick="CurtainMove('開いてる');">
    <input class="blue right-btn" type="button" value="閉" onclick="CurtainMove('閉まってる');">
    <script type="text/javascript" src="./script.js"></script>
    <br><br>
    <div class="con">
    <form action="#" method="POST" class="time-set">
      <p>時刻の設定:
      <input type="time" name="alarm_time"></p>
      <p><input class="blue item"type="submit" value="決定"></p>
    </form>
      <table class="striped tables">
        <thead>
          <tr>
              <th>日付</th>
              <th>時間</th>
          </tr>
        </thead>

        <tbody>
          <tr>
            <td>2021/6/3</td>
            <td>12:21</td>
          </tr>
          <tr>
            <td>2021/6/4</td>
            <td>10:32</td>
          </tr>
        </tbody>
      </table>
    </div>
  </body>
</html>

