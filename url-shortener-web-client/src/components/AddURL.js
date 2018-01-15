import React from 'react';
import { connect } from 'react-redux';
import { encodeURL } from '../actions';
import Button from 'react-bootstrap/lib/Button';
import FormGroup from 'react-bootstrap/lib/FormGroup';
import ControlLabel from 'react-bootstrap/lib/ControlLabel';
import FormControl from 'react-bootstrap/lib/FormControl';
import '../styles/styles.css';
import 'bootstrap/dist/css/bootstrap.css';

class AddURL extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: ''
    };

    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  render() {
    return (
      <div className="vertical-center" >
        <div className="container">
          <div className="row">
            <div className="col-md-6 col-md-offset-3">
              <h2> URL Shortener </h2>
              <form onSubmit={this.handleSubmit}>
                <FormGroup controlId="formBasicText">
                  <ControlLabel>Simplify link</ControlLabel>
                  <FormControl
                    type="text"
                    value={this.state.value}
                    placeholder="Original URL"
                    onChange={this.handleChange}
                  />
                  <div className="row">
                    <div className="col-md-8 col-md-offset-2 text-center">
                      <Button disabled={!this.state.value.length} className="btn" bsStyle="success" type="submit"> Shorten URL </Button>
                    </div>
                  </div>
                </FormGroup>
              </form>
            </div>
          </div>
        </div>
      </div>
    );
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  handleSubmit(e) {
    this.props.encodeURL(this.state.value);
    this.setState({ value: '' });
    e.preventDefault();
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    encodeURL: (origin) => dispatch(encodeURL(origin))
  };
};

export default connect(null, mapDispatchToProps)(AddURL);