<?php
  $alarm_time = $_POST['alarm_time'];
  // 設定された時刻をファイルに書き出す
  $file = 'test.txt';
  file_put_contents($file, $alarm_time);
?>
