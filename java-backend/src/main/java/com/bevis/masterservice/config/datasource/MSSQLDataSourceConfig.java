package com.bevis.masterservice.config.datasource;

import javax.sql.DataSource;

import jakarta.persistence.EntityManagerFactory;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.boot.orm.jpa.EntityManagerFactoryBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.orm.jpa.JpaTransactionManager;
import org.springframework.orm.jpa.LocalContainerEntityManagerFactoryBean;
import org.springframework.transaction.PlatformTransactionManager;
import org.springframework.transaction.annotation.EnableTransactionManagement;
import org.springframework.transaction.annotation.Transactional;

@Configuration
@EnableTransactionManagement
public class MSSQLDataSourceConfig {

    @Transactional
    @Bean(name = "MSSQLDataSource")
    @ConfigurationProperties(prefix = "spring.datasource.sqlserver")
    public DataSource mssqlDataSource() {
        return DataSourceBuilder.create().build();
    }

    @Bean
    public LocalContainerEntityManagerFactoryBean mssqlEntityManagerFactory(
            EntityManagerFactoryBuilder builder, @Qualifier("MSSQLDataSource") DataSource mssqlDataSource) {
        return builder.dataSource(mssqlDataSource)
                .packages("com.example.entity.mssql")
                .persistenceUnit("mssql")
                .build();
    }

    @Bean
    public PlatformTransactionManager mssqlTransactionManager(EntityManagerFactory mssqlEntityManagerFactory) {
        return new JpaTransactionManager(mssqlEntityManagerFactory);
    }
}
