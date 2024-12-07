package com.bevis.masterservice.entity;

import java.io.Serial;
import java.io.Serializable;
import java.util.Date;

import jakarta.persistence.Column;
import jakarta.persistence.Id;
import jakarta.persistence.MappedSuperclass;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;
import org.hibernate.annotations.UuidGenerator;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.FieldDefaults;

@Getter
@Setter
@MappedSuperclass
@FieldDefaults(level = AccessLevel.PRIVATE)
public abstract class BaseEntity implements Serializable {

    @Serial
    private static final long serialVersionUID = 596960242602342438L;

    @Id
    @UuidGenerator
    String id;

    @Column(columnDefinition = "VARCHAR(50)")
    String createdBy;

    @Column(columnDefinition = "VARCHAR(50)")
    String lastModifiedBy;

    @CreationTimestamp
    @Column(updatable = false)
    Date createdAt;

    @UpdateTimestamp
    @Column(insertable = false)
    Date updatedAt;
}
