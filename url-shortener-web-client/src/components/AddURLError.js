import React from 'react';
import Alert from 'react-bootstrap/lib/Alert';

class AddURLError extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
        <Alert bsStyle="danger">
          {this.props.errors.map((err, i) => {
            return (
              <strong key={i}> {err.message} </strong>
            );
          })}
        </Alert>
    );
  };
}

export default AddURLError;