#アラーム時刻を取得し変数へ格納
alarm_time="$1 $2 * * * ./hoge.py"
#""で囲まないと*の部分でディレクトリ内のファイルが全て出力される
echo "$alarm_time" > tmp.txt
crontab tmp.txt

#テスト用
#crontab -l
#crontab -r
