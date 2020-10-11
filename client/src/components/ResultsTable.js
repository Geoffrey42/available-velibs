import React, { useEffect, useState } from 'react';
import { Header, Table } from 'semantic-ui-react';
import axios from 'axios';

const url = "http://localhost:4242/api/fetch"

function ResultsTable() {
    const [records, setRecords] = useState([]);
    const [hits, setHits] = useState(0);
    const [total, setTotal] = useState(0);

    useEffect(() => {
        const interval = setInterval(async () => {
            try {
                const response = await axios(url);

                setRecords(response.data.records);
                setHits(response.data.nhits);
                setTotal(response.data.Total);

            } catch(err) {
                console.error("Could not fetch backend server: ", err);
            }
        }, 1000);
        return () => clearInterval(interval);
    }, [records, hits, total]);

    return (
        <Table>
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell>
                        There are {total} available velibs [YYY] meters around Splio's HQ at the moment
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