db.createUser(
    {
        user: "cinsear-user",
        pwd: "cinsear-pwd",
        roles: [
            {
                role: "readWrite",
                db: "cinsear"
            }
        ]
    }
);