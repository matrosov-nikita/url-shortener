import React from 'react';
import { connect } from 'react-redux';
import AddURL from './AddURL';
import URLModal from './URLModal';
import AddURLError from './AddURLError';

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        {this.props.errors.length > 0 &&
        <AddURLError errors={this.props.errors} />
        }
        <AddURL />
        <URLModal />
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    errors: state.errors,
  };
};


export default connect(mapStateToProps, null)(App);