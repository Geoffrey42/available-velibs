import React from 'react';
import { Header, Segment } from 'semantic-ui-react';

function AppHeader() {
    return (
        <Segment inverted>
            <Header as='h2' inverted color='teal' textAlign='left'>
                Available-velibs
            </Header>
        </Segment>
    );
}

export default AppHeader;
