package com.netflix.conductor.sql;

import com.netflix.conductor.core.config.Configuration;

import java.util.Optional;
import java.util.concurrent.TimeUnit;

public interface SQLConfiguration extends Configuration {

    String JDBC_URL_PROPERTY_NAME = "jdbc.url";
    String JDBC_URL_DEFAULT_VALUE = "";

    String JDBC_USER_NAME_PROPERTY_NAME = "jdbc.username";
    String JDBC_USER_NAME_DEFAULT_VALUE = "";

    String JDBC_PASSWORD_PROPERTY_NAME = "jdbc.password";
    String JDBC_PASSWORD_DEFAULT_VALUE = "";

    String FLYWAY_ENABLED_PROPERTY_NAME = "flyway.enabled";
    boolean FLYWAY_ENABLED_DEFAULT_VALUE = true;

    String FLYWAY_TABLE_PROPERTY_NAME = "flyway.table";
    Optional<String> FLYWAY_TABLE_DEFAULT_VALUE = Optional.empty();

    // The defaults are currently in line with the HikariConfig defaults, which are unfortunately private.
    String CONNECTION_POOL_MAX_SIZE_PROPERTY_NAME = "conductor.sql.connection.pool.size.max";
    int CONNECTION_POOL_MAX_SIZE_DEFAULT_VALUE = -1;

    String CONNECTION_POOL_MINIMUM_IDLE_PROPERTY_NAME = "conductor.sql.connection.pool.idle.min";
    int CONNECTION_POOL_MINIMUM_IDLE_DEFAULT_VALUE = -1;

    String CONNECTION_MAX_LIFETIME_PROPERTY_NAME = "conductor.sql.connection.lifetime.max";
    long CONNECTION_MAX_LIFETIME_DEFAULT_VALUE = TimeUnit.MINUTES.toMillis(30);

    String CONNECTION_IDLE_TIMEOUT_PROPERTY_NAME = "conductor.sql.connection.idle.timeout";
    long CONNECTION_IDLE_TIMEOUT_DEFAULT_VALUE = TimeUnit.MINUTES.toMillis(10);

    String CONNECTION_TIMEOUT_PROPERTY_NAME = "conductor.sql.connection.timeout";
    long CONNECTION_TIMEOUT_DEFAULT_VALUE = TimeUnit.SECONDS.toMillis(30);

    String ISOLATION_LEVEL_PROPERTY_NAME = "conductor.sql.transaction.isolation.level";
    String ISOLATION_LEVEL_DEFAULT_VALUE = "";

    String AUTO_COMMIT_PROPERTY_NAME = "conductor.sql.autocommit";
    // This is consistent with the current default when building the Hikari Client.
    boolean AUTO_COMMIT_DEFAULT_VALUE = false;

    default String getJdbcUrl() {
        return getProperty(JDBC_URL_PROPERTY_NAME, JDBC_URL_DEFAULT_VALUE);
    }

    default String getJdbcUserName() {
        return getProperty(JDBC_USER_NAME_PROPERTY_NAME, JDBC_USER_NAME_DEFAULT_VALUE);
    }

    default String getJdbcPassword() {
        return getProperty(JDBC_PASSWORD_PROPERTY_NAME, JDBC_PASSWORD_DEFAULT_VALUE);
    }

    default boolean isFlywayEnabled() {
        return getBoolProperty(FLYWAY_ENABLED_PROPERTY_NAME, FLYWAY_ENABLED_DEFAULT_VALUE);
    }

    default Optional<String> getFlywayTable() {
        return Optional.ofNullable(getProperty(FLYWAY_TABLE_PROPERTY_NAME, null));
    }

    default int getConnectionPoolMaxSize() {
        return getIntProperty(CONNECTION_POOL_MAX_SIZE_PROPERTY_NAME, CONNECTION_POOL_MAX_SIZE_DEFAULT_VALUE);
    }

    default int getConnectionPoolMinIdle() {
        return getIntProperty(CONNECTION_POOL_MINIMUM_IDLE_PROPERTY_NAME, CONNECTION_POOL_MINIMUM_IDLE_DEFAULT_VALUE);
    }

    default long getConnectionMaxLifetime() {
        return getLongProperty(CONNECTION_MAX_LIFETIME_PROPERTY_NAME, CONNECTION_MAX_LIFETIME_DEFAULT_VALUE);
    }

    default long getConnectionIdleTimeout() {
        return getLongProperty(CONNECTION_IDLE_TIMEOUT_PROPERTY_NAME, CONNECTION_IDLE_TIMEOUT_DEFAULT_VALUE);
    }

    default long getConnectionTimeout() {
        return getLongProperty(CONNECTION_TIMEOUT_PROPERTY_NAME, CONNECTION_TIMEOUT_DEFAULT_VALUE);
    }

    default String getTransactionIsolationLevel() {
        return getProperty(ISOLATION_LEVEL_PROPERTY_NAME, ISOLATION_LEVEL_DEFAULT_VALUE);
    }

    default boolean isAutoCommit() {
        return getBoolProperty(AUTO_COMMIT_PROPERTY_NAME, AUTO_COMMIT_DEFAULT_VALUE);
    }
}
