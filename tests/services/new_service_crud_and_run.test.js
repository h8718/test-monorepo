import request from 'supertest'
import retry from "jest-retries";
import common from "../common";

const testNamespace = "test-services"

describe('Test services crud operations', () => {
    beforeAll(common.helpers.deleteAllNamespaces)

    common.helpers.itShouldCreateNamespace(it, expect, testNamespace)

    common.helpers.itShouldCreateFile(it, expect, testNamespace,
        "/s1.yaml", `
direktiv_api: service/v1
image: redis
cmd: redis-server
scale: 1
`)

    common.helpers.itShouldCreateFile(it, expect, testNamespace,
        "/s2.yaml", `
direktiv_api: service/v1
image: redis
cmd: redis-server
scale: 2
`)

    let listRes;
    retry(`should list all services`, 10, async () => {
        await sleep(500)
        listRes = await request(common.config.getDirektivHost())
            .get(`/api/v2/namespaces/${testNamespace}/services`)
        expect(listRes.statusCode).toEqual(200)
        expect(listRes.body).toMatchObject({
            data: [
                {
                    "cmd": "redis-server",
                    "error": null,
                    "filePath": "/s1.yaml",
                    "id": "test-services-s1-yaml-466337cb33",
                    "image": "redis",
                    "namespace": "test-services",
                    "scale": 1,
                    "size": "medium",
                    "type": "namespace-service",
                },
                {
                    "cmd": "redis-server",
                    "error": null,
                    "filePath": "/s2.yaml",
                    "id": "test-services-s2-yaml-d396514862",
                    "image": "redis",
                    "namespace": "test-services",
                    "scale": 2,
                    "size": "medium",
                    "type": "namespace-service",
                }
            ]
        })
    })

    retry(`should list all service pods`, 10, async () => {
        await sleep(500)

        let sID = listRes.body.data[0].id
        let res = await request(common.config.getDirektivHost())
            .get(`/api/v2/namespaces/${testNamespace}/services/${sID}/pods`)
        expect(res.statusCode).toEqual(200)
        expect(res.body).toMatchObject({
            data: [
                {id: expect.stringMatching(`^${sID}(_|-)`)},
            ]
        })

        sID = listRes.body.data[1].id
        res = await request(common.config.getDirektivHost())
            .get(`/api/v2/namespaces/${testNamespace}/services/${sID}/pods`)
        expect(res.statusCode).toEqual(200)
        expect(res.body).toMatchObject({
            data: [
                {id: expect.stringMatching(`^${sID}(_|-)`)},
                {id: expect.stringMatching(`^${sID}(_|-)`)},
            ]
        })
    })

    retry(`should list all services`, 10, async () => {
        await sleep(500)

        const res = await request(common.config.getDirektivHost())
            .get(`/api/v2/namespaces/${testNamespace}/services`)
        expect(res.statusCode).toEqual(200)
        expect(res.body).toMatchObject({
            data: [
                {
                    "cmd": "redis-server",
                    "conditions": expect.arrayContaining([expect.anything()]),
                    "error": null,
                    "filePath": "/s1.yaml",
                    "id": "test-services-s1-yaml-466337cb33",
                    "image": "redis",
                    "namespace": "test-services",
                    "scale": 1,
                    "size": "medium",
                    "type": "namespace-service",
                },
                {
                    "cmd": "redis-server",
                    "conditions": expect.arrayContaining([expect.anything()]),
                    "error": null,
                    "filePath": "/s2.yaml",
                    "id": "test-services-s2-yaml-d396514862",
                    "image": "redis",
                    "namespace": "test-services",
                    "scale": 2,
                    "size": "medium",
                    "type": "namespace-service",
                }
            ]
        })
    })
});

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}