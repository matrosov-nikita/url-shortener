import React from 'react';
import AddURL from './AddURL';
import URLModal from './URLModal';

class App extends React.Component {
  render() {
    return (
      <div>
        <AddURL />
        <URLModal />      
      </div>
    );
  }
}

export default App;