import React from 'react';
import { connect } from 'react-redux';
import { toggleModal } from '../actions';
import Modal from 'react-bootstrap/lib/Modal';
import Button from 'react-bootstrap/lib/Button';
import FormControl from 'react-bootstrap/lib/FormControl';
import {CopyToClipboard} from 'react-copy-to-clipboard';

const SERVER_URL = 'http://localhost:3000';
class URLModal extends React.Component {
  constructor() {
    super();
    this.state = {
      value: '',
      copied: false,
    };

    this.handleFocus = this.handleFocus.bind(this);
    this.onCopy = this.onCopy.bind(this);
  }

  render() {
    return (
      <div>
        <Modal show={this.props.visibility} onHide={this.props.toggleModal}>
          <Modal.Header closeButton>
            <Modal.Title>Copy your short URL</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <FormControl
              type="text"
              value={`${SERVER_URL}/${this.props.urls.shortURL}`}
              onFocus={this.handleFocus}
              readOnly
            />
            <CopyToClipboard text={`${SERVER_URL}/${this.props.urls.shortURL}`} onCopy={this.onCopy}>
              <Button bsStyle="success"> Copy </Button>
            </CopyToClipboard>
          </Modal.Body>
        </Modal>
      </div>
    );
  }

  handleFocus(e) {
    e.target.select();
  }

  onCopy() {
    this.setState({ copied: true });
  }
}

const mapStateToProps = (state) => {
  return {
    urls: state.urls,
    visibility: state.modalWindowVisibility
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    toggleModal: () => dispatch(toggleModal())
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(URLModal);