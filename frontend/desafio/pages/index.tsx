import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import React, { Component } from 'react';
import PageWithJSbasedForm from "./transactions"
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import 'bootstrap/dist/css/bootstrap.css';

class Index extends Component {
  state = {};
  render() {
    return <PageWithJSbasedForm />;
  }
}

export default Index;

