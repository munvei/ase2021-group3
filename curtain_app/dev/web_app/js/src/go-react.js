import React from "react";
import ReactDOM from "react-dom"

class Layout extends React.Component {
  constructor(props) {
    super(props);

    const loc = window.location;
    const wsUrl = (loc.protocol === "https:" ? "wss:" : "ws:") + "//" + loc.host + "/test/ws";
    const ws = new WebSocket(wsUrl);
    ws.onopen = () => {
      console.log("Connected");
    }

    this.state = {
      time: "",
      isOpen: false,
      buttonMsg: "開ける",
      statusMsg: "閉まっています",
      wakeUpTime: "--：--",
      inputTime: "",
      tableItems: [],
      ws: ws,
      // url: "http://localhost",
      url: "http://54.173.221.236",
    };

    this.getDB();
    this.getSettingTime();
  }

  getSettingTime() {
    this.state.ws.send();
    this.state.ws.onmessage = (evt) => {
      const texts = evt.data.split(" ");
      this.setState({wakeUpTime: `${texts[1]}：${texts[0]}`});
    }
  }

  getDB() {
    fetch(this.state.url+"/db").then((response) => response.json()).then((responseJson) => {
      this.setState({tableItems: responseJson.rows});
    });
  }

  handleClick() {
    if (this.state.isOpen) {  // 開いている状態，----- 閉める -----
      this.setState({
        isOpen: false,
        buttonMsg: "開ける",
        statusMsg: "閉まっています",
      });
      this.state.ws.send("call: python3 alarmsys.py close");
    } else {                  // 閉まっている状態，----- 開ける -----
      this.setState({
        isOpen: true,
        buttonMsg: "閉める",
        statusMsg: "開いています",
      });
      this.state.ws.send("call: python3 alarmsys.py open_only");
    }
  }

  handleSet() {
    const splitTexts = this.state.inputTime.split(":"); 
    this.setState({wakeUpTime: this.state.inputTime});
    this.state.ws.send(`call: ./cron_make.sh ${splitTexts[1]} ${splitTexts[0]} open`);
  }
  
  handleReset() {
    this.state.ws.send("call: crontab -r");
    this.setState({wakeUpTime: "--：--"});
  }

  handleReload() {
    this.getDB();
  }

  handleTimeInput(event) {
    this.setState({inputTime: event.target.value});
  }

  componentDidMount() {
    this.timerID = setInterval(
      () => this.tick(),
      1000
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID);
  }

  tick() {
    this.setState({
      time: (new Date()).toLocaleTimeString(),
    });
  }

  render() {
    return (
      <div>
        {/* header */}
        <div id="header">
          <header className="blue white-text">
            <h1 className="center">ACOS 管理ページ</h1>
          </header>
        </div>

        {/* current time */}
        <center id="TimeArea"><font size="6">{ this.state.time }</font></center>

        <table><tbody>
        {/* カーテン操作用ボタン */}
          <tr>
            <th width="150">現在の状態：</th>
            <th colSpan="2">{ this.state.statusMsg }</th>
            <th colSpan="2"><center><button onClick={ () => {this.handleClick()} }>{ this.state.buttonMsg }</button></center></th>
          </tr>

        {/* 起床時刻の設定 */}
          <tr>
            <th width="150">現在の設定時間：</th>
            <th width="150">{ this.state.wakeUpTime }</th>
            <th width="150"><button onClick={ () => {this.handleReset()} }>リセット</button></th>
            <th width="150"><input type="time" name="time" onChange={ (event) => {this.handleTimeInput(event)} } /></th>
            <th width="150"><button onClick={ () => {this.handleSet()} }>設定</button></th>
          </tr>
        </tbody></table>

        <br/><br/>

        {/* 起床ログを表示 */}
        <div id="table">
          <h6>
            起床ログ　
            <button onClick={ () => {this.handleReload()} }>更新</button>
          </h6>
          <table><tbody>
            <tr>
              <th><label>ID</label></th>
              <th><label>DATE</label></th>
              <th><label>MESSAGE</label></th>
            </tr>
            { this.state.tableItems.map((item) => {
              return(
                <tr>
                  <td>{ item.Id }</td>
                  <td>{ item.Date }</td>
                  <td>{ item.Msg }</td>
                </tr>
              );
            }) }
          </tbody></table>
        </div>
      </div>
    );
  }
}

const app = document.getElementById("output");
ReactDOM.render(<Layout />, app);
