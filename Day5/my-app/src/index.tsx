import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import {BrowserRouter, Switch, Route} from "react-router-dom";
import Login from "./components/login";
import Home from "./components/home";
import Message from "./components/message";

ReactDOM.render(
    <React.StrictMode>
        <BrowserRouter>
            <Switch>
                <Route exact path="/">
                    <Login/>
                </Route>
            </Switch>
        </BrowserRouter>
        <BrowserRouter>
            <Switch>
                <Route exact path="/room">
                    <Home/>
                </Route>
            </Switch>
        </BrowserRouter>
    </React.StrictMode>,
    document.getElementById("root")
);