import './App.css';
import React from 'react'

const URL = [
    "localhost:8082",
    "localhost:8083/api/user?name=",
    "localhost:8081/count"
];
  
class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = { indexProgram: 0, indexRequest: 0, url: URL[0], inputText: "", result: "", requestStatus: "" };

    }

    handleChangeProgram(event) {
        this.setState({ indexProgram: event.target.value, url: URL[event.target.value] });
        this.clearFields();
    }

    handleChangeRequst(event) {
        this.setState({ indexRequest: event.target.value });
        this.clearFields();
    }

    clearFields() {
        this.setState({ inputText: "", requestStatus: "", result: "" })
    }

    handleChangeInput(event) {
        this.setState({ inputText: event.target.value })
        if (this.state.indexProgram == 1) {
            this.setState({ url: URL[1] + encodeURIComponent(event.target.value) })
        }
    }
    handleClick() {
        this.setState({ requestStatus: "", result: "" })
        if (this.state.indexRequest == 0) {
            fetch(`http://${this.state.url}`).then(async (response) => {
                let responseValue = await response.text();
                let responseStatus = response.status;
                this.setState({ requestStatus: responseStatus, result: responseValue });
            }).catch((err) => {
                this.setState({ requestStatus: 500, result: "Ошибка соединения!" });
            })
        } else {
            let formData = new FormData();
            formData.append('count', this.state.inputText);
            fetch(`http://${this.state.url}`, {
                method: 'POST',
                body: formData
            }).then(async (response) => {
                let responseValue = await response.text();
                let responseStatus = response.status;
                this.setState({ requestStatus: responseStatus, result: responseValue });
            }).catch((err) => {
                this.setState({ requestStatus: 500, result: "Ошибка соединения!" });
            })
        }
    }
    render() {
        return (
            <div className='App'>
                <div className='head'>
                    <div className='head_program'>
                        <label className='text'>Программа</label>
                        <select className='select' value={this.state.indexProgram} onChange={(event) => this.handleChangeProgram(event)}>
                            <option value="0">Hello</option>
                            <option value="1">Query</option>
                            <option value="2">Count</option>
                        </select>
                    </div>
                    <button className='button_run' onClick={() => this.handleClick()}>Отправить</button>
                    <div className='head_request'>
                        <label className='text'>Запрос</label>
                        <select className='select' value={this.state.indexRequest} onChange={(event) => this.handleChangeRequst(event)}>
                            <option value="0">GET</option>
                            {this.state.indexProgram == 2 && <option value="1">POST</option>}
                        </select>
                    </div>
                </div>
                <hr className='separator'></hr>
                {(this.state.indexProgram == 1 || (this.state.indexProgram == 2 && this.state.indexRequest == 1)) &&
                    <>
                        <div className='body'>
                            {this.state.indexProgram == 1 && <>
                                <div className='text'><strong>Query Params</strong></div>
                                <label className='text'>name: <input className="input" value={this.state.inputText} onChange={(event) => this.handleChangeInput(event)} /></label>
                            </>}
                            {this.state.indexProgram == 2 && <>
                                <div className='text'><strong>Form Params</strong></div>
                                <label className='text'>count: <input className="input" value={this.state.inputText} onChange={(event) => this.handleChangeInput(event)} /></label>
                            </>}
                        </div>
                        <hr className='separator'></hr>
                    </>}
                <div className='body'>
                    <label className='text'><strong>URL: </strong> {this.state.url}</label>
                    <div className='text'><strong>Результат: </strong>{this.state.result}</div>
                    <div className='text'>Статус: {this.state.requestStatus}</div>
                </div>
            </div>
        );
    }
}

export default App;