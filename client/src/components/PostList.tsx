import "../App.css";

import React from "react";
import { Link } from "react-router-dom";

const BasicPostList: React.FC = () => {
    return (
        <div style={{ width: "25vw", margin: "auto", textAlign: "center" }}>
            <h4>{"Post List"}</h4>
            <ul>
                <li>
                    <Link to="/post/1">{"Inspirational Quotes"}</Link>
                    {" by Aiken"}
                </li>
            </ul>
        </div>
    );
};

export default BasicPostList;
