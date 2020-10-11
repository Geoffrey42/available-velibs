import React from 'react';
import { Header, Table } from 'semantic-ui-react';

function ResultsTable() {
    return (
        <Table>
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell>
                        There are [XXX] available velibs [YYY] meters around Splio's HQ at the moment
                    </Table.HeaderCell>
                </Table.Row>
            </Table.Header>
            <Table.Body>
                <Table.Row>
                    <Table.Cell>
                        <Header as='h4'>
                            <Header.Content>
                                [STATION NAME] station
                                <Header.Subheader>
                                    Available velibs: [NUMBER]
                                </Header.Subheader>
                            </Header.Content>
                        </Header>
                    </Table.Cell>
                </Table.Row>
            </Table.Body>
        </Table>
    );
}

export default ResultsTable;