import React from 'react';
import Alert from 'react-bootstrap/lib/Alert';

class AddURLError extends React.Component {
  render() {
    return (
        <Alert bsStyle="danger">
          {this.props.errors.map((err, i) => {
            return (
              <strong key={i}> {err.detail || 'Something went wrong'} </strong>
            );
          })}
        </Alert>
    );
  };
}

export default AddURLError;