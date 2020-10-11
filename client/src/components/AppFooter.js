import React from 'react';
import { Header, Segment, Image } from 'semantic-ui-react';

function AppFooter() {
    return (
        <Segment inverted>
            <Header as='h5' inverted color='teal' textAlign='left'>
                Made by Geoffrey42 with <Image src='/public/red_heart_00.png' />
                <br />
                <Header.Subheader>
                    This application is part of Splio's homework technical test
                </Header.Subheader>
            </Header>
        </Segment>
    );
}

export default AppFooter;
