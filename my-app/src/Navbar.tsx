import React from 'react';
import {Link} from "react-router-dom";
import {MenuItem, Button, Toolbar, AppBar} from "@mui/material";
import "./App.css"


function Navbar() {
    const pages = [
        {
            name: "Home",
            endpoint: "/"
        },
        {
            name: "Form",
            endpoint: "/form"
        },
    ]

    return (
        <AppBar position="static">
            <Toolbar>
                {pages.map((page) => (
                    <Link
                        color="inherit"
                        to={page.endpoint}
                        style={{textDecoration: 'None', color: "white"}}
                    >
                        <Button
                            color="inherit"
                            key={page.name}
                        >
                            {page.name}
                        </Button>
                    </Link>
                ))}
            </Toolbar>
        </AppBar>
    )
}

export default Navbar;