import './App.css';
import Logo from "./assets/logo.png"
import React, { useEffect, useState } from "react"
import './index.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import {
  Route,
  HashRouter,
  Navigate
} from "react-router-dom";
import { Nav, Button, Modal } from "react-bootstrap"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {
  faRightFromBracket,
  faExchange,
  faBook,
  faCheck,
  faLock,
} from '@fortawesome/free-solid-svg-icons'
import { Routes } from "react-router-dom";
import Login from "./components/Login";
import Pencairan from "./components/Pencairan";
import Laporan from "./components/Laporan";
import ChangePassword from "./components/ChangePassword";
import Popup from 'reactjs-popup';
import axios from "axios";

function App() {
  const [isLoggedIn, setLogin] = useState()
  const [userName, setUserName] = useState("")
  const [isCollapsed, setIsCollapsed] = useState(true)
  const [activeMenu, setActiveMenu] = useState("")
  const [show, setShow] = useState(false)

  const handleClose = () => setShow(false)
  const handleShow = () => setShow(true)

  const handleLogOut = () => {
    localStorage.removeItem("info")
    localStorage.removeItem("user_id")
    window.location.href = "/"
  }

  const getUsername = () => {
    console.log("getUsername")
    axios.get("http://localhost:8080/findUser?userId=" + localStorage.getItem("user_id"))
      .then((response) => {
        setUserName(response.data.name)
      }).catch(function (error) {
        console.log(error)
        setUserName("please reload this site")
      })
  }

  const collapseNavbar = () => {
    isCollapsed ? setIsCollapsed(false) : setIsCollapsed(true)
    const navTrigger = document.getElementById("nav-trigger")
    const triggerDiv = document.getElementsByClassName("eventCollapse")
    const triggerSvg = document.getElementsByClassName("eventCollapse2")

    if (isCollapsed) {
      navTrigger.style.width = "110px"

      for (let i = 0; i < triggerDiv.length; i += 1) {
        triggerDiv[i].style.display = 'none';
      }

      for (let i = 0; i < triggerSvg.length; i += 1) {
        triggerSvg[i].style.paddingRight = "0"
        triggerSvg[i].style.textAlign = "center"
        triggerSvg[i].style.width = "100%"
      }

    } else {
      navTrigger.style.width = "350px"


      for (let i = 0; i < triggerSvg.length; i += 1) {
        triggerSvg[i].style.removeProperty("width")
        triggerSvg[i].style.paddingRight = "1rem"
        triggerSvg[i].style.removeProperty("textAlign")
      }

      for (let i = 0; i < triggerDiv.length; i += 1) {
        triggerDiv[i].style.display = "inline-block"
      }
    }
  }

  useEffect(() => {
    localStorage.getItem("info") != null ? setLogin(true) : setLogin(false)
    getUsername()
  }, [])

  if (!isLoggedIn) {
    return (
      <>
        <HashRouter>
          <Routes>
            <Route path="/" element={<Login />} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        </HashRouter>
      </>
    )
  }

  else {
    return (
      <>
        <HashRouter>
          <div className="d-flex">
            <div>
              <Nav className="flex-column" id="nav-trigger">
                <div className="d-flex align-items-center menubar-brand">
                  <img id="image-trigger" onClick={() => collapseNavbar()} src={Logo} alt={"Logo BSIM"}></img>
                  <div className="d-flex flex-column">
                    <h3 className="eventCollapse text-white">Pencairan Kredit</h3>
                    <h6 className="eventCollapse text-white">Hello, {userName}</h6>
                  </div>
                </div>
                <div className="menubar-list">
                  <Popup
                    trigger={
                      <Nav.Link
                        className={activeMenu == "transaksi" ? "nav-active" : ""}>
                        <FontAwesomeIcon className="eventCollapse2" icon={faExchange} /><span className="eventCollapse">Transaksi</span>
                      </Nav.Link>
                    }
                    position="right center"
                    on="click"
                  >
                    <div className='navLinkPopup'>
                      <Nav.Link
                        href="#/Pencairan"
                        onClick={() => setActiveMenu("transaksi")} >
                        <FontAwesomeIcon className="eventCollapse2" icon={faCheck} /><span className="eventCollapse">   Checklist Pencairan</span>
                      </Nav.Link>
                    </div>
                  </Popup>
                  <Nav.Link
                    href="#/Laporan"
                    onClick={() => setActiveMenu("laporan")}
                    className={activeMenu == "laporan" ? "nav-active" : ""}>
                    <FontAwesomeIcon className="eventCollapse2" icon={faBook} /><span className="eventCollapse">Laporan</span>
                  </Nav.Link>
                  <hr className="text-white"></hr>
                  <Nav.Link
                    href="#/ChangePassword"
                    onClick={() => setActiveMenu("changePassword")}
                    className={activeMenu == "changePassword" ? "nav-active" : ""}>
                    <FontAwesomeIcon className="eventCollapse2" icon={faLock} /><span className="eventCollapse">Ubah Password</span>
                  </Nav.Link>
                  <Nav.Link
                    onClick={() => handleShow()}>
                    <FontAwesomeIcon className="eventCollapse2" icon={faRightFromBracket} /><span className="eventCollapse">Keluar</span>
                  </Nav.Link>
                  <hr></hr>
                </div>
              </Nav>
            </div>
            <div className="content">
              <Routes>
                <Route path="/" />
                <Route path="/Pencairan" element={<Pencairan />} />
                <Route path="/Laporan" element={<Laporan />} />
                <Route path="/ChangePassword" element={<ChangePassword />} />
                <Route path="*" element={<Pencairan />} />
              </Routes>
            </div>
          </div>
          <Modal centered show={show} onHide={handleClose}>
            <Modal.Header closeButton className="backgroundRed">
              <Modal.Title className="text-white">Konfirmasi</Modal.Title>
            </Modal.Header>
            <Modal.Body>Apakah anda yakin mau keluar?</Modal.Body>
            <Modal.Footer>
              <Button variant="secondary" className="backgroundBlack" onClick={handleClose}>
                Tidak
              </Button>
              <Button variant="primary" className="backgroundRed" onClick={handleLogOut}>
                Ya
              </Button>
            </Modal.Footer>
          </Modal>
        </HashRouter>
      </>
    );
  }
}

export default App;
