import React from 'react';
import { Header, Segment, Image } from 'semantic-ui-react';
import heart from '../red_heart_01.png';
function AppFooter() {
    return (
        <Segment inverted style={{marginBottom: 0}}>
            <Header as='h5' inverted color='teal' textAlign='left'>
                Made by Geoffrey42 with <Image src={heart} size={'tiny'} />
                <br />
                <Header.Subheader>
                    This application is part of Splio's homework technical test
                </Header.Subheader>
            </Header>
        </Segment>
    );
}

export default AppFooter;
